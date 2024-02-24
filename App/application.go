package App

import (
	"TigerPopulation/Controllers/Middleware"
	"TigerPopulation/Utils/dbConfig"
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	router = gin.New()
)

func StartApplication() {
	//defer ants.Release()
	defer glog.Flush()
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logpathsskipped := []string{"/healthz", "/readyz"}
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: logpathsskipped}))

	devConfig := dbConfig.ViperConfigDev()
	//common.InitPaseto()
	SetupRouteUtils()
	router.Use(cors.Default())
	router.Use(Middleware.SetSessionToken())
	router.Use(Middleware.CustomTokenMiddleware(devConfig.JwtSecretKey))
	router.Use(gzip.Gzip(gzip.BestSpeed))
	MapDevUrls()
	MapLiveUrls()
	//	imjobs.Setupjobs()

	srv := &http.Server{
		Addr:    devConfig.Port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	glog.V(2).Infoln("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Errorln("Server forced to shutdown: ", err)
	}

	glog.V(2).Infoln("Server exiting")
}
