package conf

import ()

const (
	// Release
	RELEASE_MODE = false

	// Project
	PROJECT_NAME = "Echo Web"

	SERVER_ADDR = ":5178"

	// Domains
	DOMAIN_API    = "api.men.n.mi.com"
	DOMAIN_WWW    = "men.n.mi.com"
	DOMAIN_SOCKET = "socket.api.men.n.mi.com"

	// Session
	SESSION_STORE = REDIS // REDIS,FILE,COOKIE(COOKIE Default)

	// Cache
	CACHE_STORE = REDIS //REDIS,MEMCACHED,IN_MEMORY(IN_MEMORY Default)

	// Tmpl
	TMPL_TYPE   = PONGO2   // PONGO2,TEMPLATE(TEMPLATE Default)
	TMPL_DATA   = BINDATA       // BINDATA,FILE(FILE Default)
	TMPL_DIR    = "template/pongo2" // PONGO2(template/pongo2),TEMPLATE(template)
	TMPL_SUFFIX = ".html"    // .html,.tpl

	// Static
	STATIC_TYPE = BINDATA // BINDATA,FILE(FILE Default)

	// MySQL
	DB_NAME      = "xmstore-act"
	DB_USER_NAME = "root"
	DB_PASSWORD  = ""
	DB_HOST      = "127.0.0.1"
	DB_PORT      = "3306"

	// Redis
	REDIS_SERVER = "127.0.0.1:6379"
	REDIS_PWD    = ""

	// Memcached
	MEMCACHED_SERVER = "localhost:11211"
)

const (
	// Template Type
	PONGO2   = "pongo2"
	TEMPLATE = "template"

	// Bindata
	BINDATA = "bindata"

	// File
	FILE = "file"

	// Redis
	REDIS = "redis"

	// Memcached
	MEMCACHED = "memcached"

	// Cookie
	COOKIE = "cookie"

	// In Memory
	IN_MEMORY = "InMemory"
)
