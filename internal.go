package gologger

// LoggerType is used to indicate the the logger type.
// Where or How to Log.
type LoggerType string

const (
	// Glog will use "github.com/golang/glog" to log this will use the leveled
	// handled log for logging like glog.V().Infoln().
	// This will not apply the category filter provided by gologger. By Providing
	// `gologger.EnableC` as options, can apply the category filter before log
	// to glog.
	Glog          LoggerType = "glog"

	// Console Logger. Simple Will use the fmt package. Two modes available for
	// Logging. Colored Log or Plain text log to Console. Provides Options as
	// `gologger.Plain` and `gologger.Colored`
	Console                  = "console"

	// All logs to filesystem. The file path as options Needs to be provided.
	// if path is not provided logs to the default package location.
	File                     = "file"

	// Stores logs in elasticsearch in logstash format. Kibana is able
	// to use those logs. The elasticsearch location needs to be set via options.
	// If no options is set then will use `localhost:9200` to communicate with
	// elasticsearch.
	ElasticSearch            = "elasticsearch"
)

const (
	EnableC  = iota
	Plain
	Colored
)


var activeLogger map[string]*Logger

func (t LoggerType) String() string {
	return string(t)
}

func init() {
	activeLogger = make(map[string]*Logger)
}

type config struct {
}
