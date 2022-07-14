local branches = ["dev","testnet", "main", "dev-local", "dev-15"];
local branch = std.extVar("build.branch");
local projectName = "svc-broadcaster";
local imageName = "symbiosisfinance/%s" % projectName;

local projectEnv(branch) = (
    if branch == "dev" then
		"dev"
    else if branch == "dev-local" then
		"dev-local"
    else if branch == "dev-15" then
		"dev-15"
	else if branch == "testnet" then
        "testnet"
	else if branch == "main" then
        "mainnet"
    else
        error "No config for branch %s" % branch
);

local imageNameWithTag(branch) = imageName + ":" + projectEnv(branch);

local projectNameWithEnv(branch) = projectName + "-" + projectEnv(branch);

local composeProjectName(branch) =
	"--project-name "+projectName+"-"+projectEnv(branch);

local dockerCompose(branch) = (if branch == "dev" then
        "docker-compose.dev.yml"
	else if branch == "dev-local" then
        "docker-compose.dev-local.yml"
	else if branch == "dev-15" then
        "docker-compose.dev-15.yml"
	else if branch == "testnet" then
        "docker-compose.testnet.yml"
	else if branch == "main" then
        "docker-compose.main.yml"
    else
        error "No project name for branch %s" % branch);

local domain(branch) = (
	if branch == "dev" then
        "api.dev"
    else if branch == "dev-15" then
        "api.dev-15"
    else if branch == "dev-local" then
        "api.dev-local"
    else if branch == "testnet" then
        "api.testnet"
    else if branch == "main" then
        "api"
    else
        error "No domain for branch %s" % branch
) + ".symbiosis.finance";

local deployNode(branch) = (
	if  branch == "dev" || branch == "testnet" || branch == "dev-15" || branch == "dev-local" then
		{
			"role": "dev"
		}
    else if branch == "main" then
		{
			"role": "prod"
		}
    else
        error "No node for branch %s" % branch
);

local chainKeys(branch) = (
    if branch == "dev" || branch == "dev-15" || branch == "dev-local" || branch == "testnet" then
        {
            "BROADCASTER_CHAIN_256_KEY": {
                 "from_secret": "chain_256_HECO_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_43113_KEY": {
                 "from_secret": "chain_43113_AVAX_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_4_KEY": {
                 "from_secret": "chain_4_ETH_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_65_KEY": {
                 "from_secret": "chain_65_OEC_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_80001_KEY": {
                 "from_secret": "chain_80001_POLYGON_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_97_KEY": {
                 "from_secret": "chain_97_BSC_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_28_KEY": {
                 "from_secret": "chain_28_BOBA_TESTNET_dev"
            },
            "BROADCASTER_CHAIN_200101_KEY": {
                 "from_secret": "chain_200101_MILKOMEDA_TESTNET_dev"
            }
        }
    else if branch == "main" then
        {
            "BROADCASTER_CHAIN_137_KEY": {
                "from_secret": "chain_137_POLYGON_MAINNET_mainnet"
            },
            "BROADCASTER_CHAIN_1_KEY": {
                "from_secret": "chain_1_ETH_MAINNET_mainnet"
            },
            "BROADCASTER_CHAIN_288_KEY": {
                "from_secret": "chain_288_BOBA_MAINNET_mainnet"
            },
            "BROADCASTER_CHAIN_43114_KEY": {
                "from_secret": "chain_43114_AVAX_MAINNET_mainnet"
            },
            "BROADCASTER_CHAIN_56_KEY": {
                "from_secret": "chain_56_BSC_MAINNET_mainnet"
            },
            "BROADCASTER_CHAIN_2001_KEY": {
                "from_secret": "chain_2001_MILKOMEDA_MAINNET_mainnet"
            }
        }
    else
        error "No private keys for branch %s" % branch
);

[{
     "kind": "pipeline",
     "type": "docker",
     "name": "build",
     "steps": [
         {
             "name": "test",
             "image": "golang:1.17.5",
             "commands": [
                "go test -v ./... "
             ]
         },{
             "name": "image",
             "image": "plugins/docker",
              "settings" : {
                 "repo": imageName,
                 "build_args": [
                     "CONFIG=%s" % projectEnv(branch),
                     "VERSION=git-${DRONE_COMMIT_SHA:0:8}"
                 ],
                 "dockerfile": "Dockerfile",
                 "auto_tag": false,
                 "insecure": false,
                 "purge": false,
                 "username": {
                   "from_secret": "docker_username"
                 },
                 "password": {
                     "from_secret": "docker_password"
                 },
                 "tags": [projectEnv(branch)]
             },
             "when": {
                 "branch": branches,
                 "event": "push"
             }
         }
     ]
},{
     "kind": "pipeline",
     "type": "docker",
     "name": "deploy",
     "steps": [{
             "name": "docker-compose",
             "image": "docker/compose",
             "environment": {
                    "DEPLOY_IMAGE": imageNameWithTag(branch),
                    "DEPLOY_PROJECT": projectNameWithEnv(branch),
                    "DEPLOY_DOMAIN": domain(branch),
                    "DEPLOY_PROJECT_ENV": projectEnv(branch),
                    "DEPLOY_PATH_PREFIX": "/broadcaster/",
                    "POSTGRES_DB": {
                        "from_secret": "postgres_database"
                    },
                    "POSTGRES_USER": {
                        "from_secret": "postgres_user"
                    },
                    "POSTGRES_PASSWORD": {
                        "from_secret": "postgres_password"
                    },
             } + chainKeys(branch),
             "commands": [
                 "docker pull symbiosisfinance/" + projectName + ":" + projectEnv(branch),
                 "docker-compose " + composeProjectName(branch) + " -f docker-compose.yml -f " + dockerCompose(branch) + " up -d",
             ],
             "volumes": [
                 {
                     "name": "dockersock",
                     "path": "/var/run/docker.sock"
                 },
                 {
                     "name": "dockerconfig",
                     "path": "/root/.docker/config.json"
                 }
             ],
             "when": {
                 "branch": branches,
                 "event": "push"
             }
         }
     ],
     "volumes": [
         {
             "name": "dockersock",
             "host": {
                 "path": "/var/run/docker.sock"
             }
         }, {
             "name": "dockerconfig",
             "host": {
                 "path": "/root/.docker/config.json"
             }
         }
     ],
     "depends_on": [
         "build"
     ],
     "node": deployNode(branch)
 }]
