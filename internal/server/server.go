package server

//type Server struct {
//	log log.Logger
//
//	listener http.Server
//
//	config   config.Server
//	router   *web.Web
//	networks map[string]*network.Network
//
//	lock *sync.Mutex
//
//	processing sync.WaitGroup
//}

//func New(
//	log log.Logger,
//	config config.Server,
//	lock *sync.Mutex,
//	networks map[string]*network.Network,
//) (*Server, error) {
//	router := web.New(chi.NewMux())
//	server := &Server{
//		log:      log,
//		router:   router,
//		config:   config,
//		lock:     lock,
//		networks: networks,
//	}
//
//	router.Route("/v1/", func(web *web.Web) {
//		web.Get("/status", web.ServeHandler(server.handleStatus))
//
//		{
//			web := web.With(server.auth)
//			web.Post(
//				"/broadcast/{network}/{bridge}",
//				web.ServeHandler(server.handleBroadcast),
//			)
//		}
//	})
//
//	return server, nil
//}

//func (server *Server) Serve() error {
//	server.listener = http.Server{
//		Addr:    server.config.Address.ToString(),
//		Handler: server.router,
//	}
//
//	err := server.listener.ListenAndServe()
//	if err != http.ErrServerClosed {
//		return err
//	}
//
//	return nil
//}
//
//func (server *Server) Stop() error {
//	return server.listener.Shutdown(context.TODO())
//}
//
//func (server *Server) auth(next web.Handler) web.Handler {
//	return func(context *web.Context) error {
//		if server.config.Auth != nil {
//			user, pass, ok := context.GetRequest().BasicAuth()
//			facts := karma.
//				Describe("user", user).
//				Describe("pass", "<redacted>")
//
//			if !ok {
//				return context.Error(
//					http.StatusForbidden,
//					facts.Reason(
//						"auth requested but not provided by the client",
//					),
//					"authentication is required",
//				)
//			}
//
//			if user != server.config.Auth.User {
//				return context.Error(
//					http.StatusForbidden,
//					facts.Reason("auth failed: user does not match"),
//					"invalid credentials",
//				)
//			}
//
//			if pass != server.config.Auth.Pass {
//				return context.Error(
//					http.StatusForbidden,
//					facts.Reason("auth failed: pass does not match"),
//					"invalid credentials",
//				)
//			}
//		}
//
//		return next(context)
//	}
//}
//
//func (server *Server) handleStatus(context *web.Context) error {
//	return context.OK()
//}
//
//func (server *Server) handleBroadcast(context *web.Context) error {
//	var (
//		networkName   = context.GetURLParam("network")
//		bridgeAddress = context.GetURLParam("bridge")
//	)
//
//	server.processing.Add(1)
//	defer server.processing.Done()
//
//	facts := karma.
//		Describe("network", networkName).
//		Describe("bridge", bridgeAddress)
//
//	if !eth_common.IsHexAddress(bridgeAddress) {
//		return context.BadRequest(
//			facts.Reason("bridge is not an address"),
//			"given bridge is not a valid address: %s",
//			bridgeAddress,
//		)
//	}
//
//	{
//		bridgeAddress := eth_common.HexToAddress(bridgeAddress)
//
//		network := server.networks[networkName]
//		if network == nil {
//			return context.BadRequest(
//				facts.Reason("given network is not registered in the config"),
//				"unsupported network: %s",
//				network,
//			)
//		}
//
//		var request broadcaster_client_schema.RequestBroadcast
//
//		err := render.Bind(context.GetRequest(), &request)
//		if err != nil {
//			return context.BadRequest(
//				facts.Reason(err),
//				"unable to parse request body: %s",
//				err,
//			)
//		}
//
//		session, err := network.StartSession(bridgeAddress)
//		if err != nil {
//			return context.Error(
//				http.StatusBadGateway,
//				err,
//				"unable to start signing session",
//			)
//		}
//
//		err = session.ValidateSignature(request)
//		if err != nil {
//			return context.Error(
//				http.StatusBadGateway,
//				err,
//				"signature does not match",
//			)
//		}
//
//		transaction, err := session.CreateTransaction(request)
//		if err != nil {
//			return context.Error(
//				http.StatusBadGateway,
//				err,
//				"signature does not match",
//			)
//		}
//
//		err = session.SendTransaction(transaction, server.lock)
//		if err != nil {
//			return context.Error(
//				http.StatusBadGateway,
//				err,
//				"unable to send transaction",
//			)
//		}
//	}
//
//	return context.OK()
//}
