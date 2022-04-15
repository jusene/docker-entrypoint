package conf

import (
	"fmt"
	"os"
	"path/filepath"
)

var HOME, _ = filepath.Abs(".")
var ENV = os.Getenv("env")

var JAVA_OPTS = fmt.Sprintf("-server -Xms512m -Xmx512m -Xmn256m " +
	"-XX:SurvivorRatio=8 -Xss256k -XX:-UseAdaptiveSizePolicy " +
	"-XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=%s/java.hprof -XX:ErrorFile=%s/hs_err_pid.log "+
	"-XX:MaxTenuringThreshold=15 -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+UseCMSCompactAtFullCollection " +
	"-XX:+CMSParallelRemarkEnabled -XX:+CMSPermGenSweepingEnabled -XX:+UseFastAccessorMethods " +
	"-XX:+CMSClassUnloadingEnabled -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=70 " +
	"-verbose:gc -Xloggc:%s/gc.log -XX:+PrintHeapAtGC -XX:+PrintGCDetails -XX:+PrintGCDateStamps " +
	"-XX:+PrintTenuringDistribution -XX:+PrintGCApplicationStoppedTime "+
	"-XX:-UseCompressedOops -Djava.awt.headless=true -Dsun.net.client.defaultConnectTimeout=10000 -Dsun.net.client.defaultReadTimeout=30000 " +
	"-Djava.io.tmpdir=/app/tmp " +
	"-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.port=7777 -Dcom.sun.management.jmxremote.local.only=false " +
	"-Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false " +
	"-Dspring.profiles.active=%s -jar", HOME, HOME, HOME, EnvMap[ENV])


var EnvMap = map[string]string{
	"daily": "dev",
	"qa": "qa",
	"pre": "pre",
	"prod": "prod",
}

