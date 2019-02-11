module github.com/mandocaesar/go-skeleton
require github.com/mandocaesar/go-skeleton/domain/authentication v0.0.0
replace github.com/mandocaesar/go-skeleton/domain/authentication => ./domain/authentication/
require (
	github.com/gin-contrib/cors v0.0.0-20190101123304-5e7acb10687f
	github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.2.0
	github.com/hellofresh/health-go v2.0.2+incompatible
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/sebest/logrusly v0.0.0-20180315190218-3235eccb8edc
	github.com/segmentio/go-loggly v0.5.0 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/viper v1.3.1
	github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43 // indirect
	go.elastic.co/apm/module/apmgin v1.2.0
	go.elastic.co/apm/module/apmlogrus v1.2.0
	google.golang.org/grpc v1.18.0
)
