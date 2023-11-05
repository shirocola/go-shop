package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load write timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load body limit failed: %v", err)
				}
				return b
			}(),
			fileLimit: func() int {
				f, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("load file limit failed: %v", err)
				}
				return f
			}(),
			gcpbucket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			protocal: envMap["DB_PROTOCAL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnection: func() int {
				m, err := strconv.Atoi(envMap["DB_MAX_CONNECTION"])
				if err != nil {
					log.Fatalf("load max connection failed: %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpireAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRE_AT"])
				if err != nil {
					log.Fatalf("load access expire at failed: %v", err)
				}
				return t
			}(),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) Url() string {
	return fmt.Sprintf("%s:%d,", a.host, a.port)
}

func (a *app) Name() string {
	return a.name
}

func (a *app) Version() string {
	return a.version
}

func (a *app) ReadTimeout() time.Duration {
	return a.readTimeout
}

func (a *app) WriteTimeout() time.Duration {
	return a.writeTimeout
}

func (a *app) BodyLimit() int {
	return a.bodyLimit
}

func (a *app) FileLimit() int {
	return a.fileLimit
}

func (a *app) GCPBucket() string {
	return a.gcpbucket
}

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

type db struct {
	host          string
	port          int
	protocal      string
	username      string
	password      string
	database      string
	sslMode       string
	maxConnection int
}

func (c *config) Db() IDbConfig {
	return c.db
}
func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host, d.port, d.username, d.password, d.database, d.sslMode)
}

func (d *db) MaxOpenConns() int {
	return d.maxConnection
}

type IJwtConfig interface {
	AdminKey() []byte
	SecretKey() []byte
	ApiKey() []byte
	AccessExpireAt() int
	RefreshExpireAt() int
	SetJwtAccessExpire(t int)
	SetJwtRefreshExpire(t int)
}

type jwt struct {
	adminKey        string
	secretKey       string
	apiKey          string
	accessExpireAt  int //sec
	refreshExpireAt int //sec
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

func (j *jwt) AdminKey() []byte {
	return []byte(j.adminKey)
}

func (j *jwt) SecretKey() []byte {
	return []byte(j.secretKey)
}

func (j *jwt) ApiKey() []byte {
	return []byte(j.apiKey)
}

func (j *jwt) AccessExpireAt() int {
	return j.accessExpireAt
}

func (j *jwt) RefreshExpireAt() int {
	return j.refreshExpireAt
}

func (j *jwt) SetJwtAccessExpire(t int) {
	j.accessExpireAt = t
}

func (j *jwt) SetJwtRefreshExpire(t int) {
	j.refreshExpireAt = t
}
