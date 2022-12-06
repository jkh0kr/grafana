package mtctx

import (
	"context"
	"fmt"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/sqlstore/session"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type serviceImpl struct {
	clientset *kubernetes.Clientset
	cache     map[int64]*TenantInfo
	watchers  map[int64]watch.Interface
}

// POC: shows stackID of tenant in current request
func (s *serviceImpl) showTenantInfo(c *models.ReqContext) response.Response {
	t, err := TenantInfoFromContext(c.Req.Context())
	if err != nil {
		fmt.Println("Could not find tenant info", err)
	}

	return response.JSON(200, map[string]interface{}{
		"stackID": t.StackID,
	})
}

// Gets config.ini from kubernetes and returns a watcher to listen for changes
func (s *serviceImpl) GetStackConfigWatcher(ctx context.Context, stackID int64) (watch.Interface, error) {
	if s.clientset == nil {
		return nil, fmt.Errorf("missing error")
	}

	return s.clientset.CoreV1().ConfigMaps("hosted-grafana").Watch(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata[name]=%s", stackName(stackID)),
	})
}

// MIDDLEWARE: Adds TenantInfo
func (s *serviceImpl) InitializeTenantConfig(ctx context.Context) context.Context {
	// TODO
	// get user from context
	// stackID is on context

	var stackID int64 = 0000

	if info, ok := s.cache[stackID]; ok {
		return ContextWithTenantInfo(ctx, info)
	}

	// get the initial config map
	config, err := s.clientset.CoreV1().ConfigMaps("hosted-grafana").Get(context.TODO(), stackName(stackID), metav1.GetOptions{})
	if err != nil {
		fmt.Println("Error getting config map:", err)
	}

	ti := buildTenantInfoFromConfig(stackID, config)
	ctx = ContextWithTenantInfo(ctx, ti)

	// Get config watcher
	w, err := s.GetStackConfigWatcher(context.TODO(), stackID)
	if err != nil {
		fmt.Println("Error getting watcher for stackID:", stackID)
	}

	// TODO should we check to see if we already have a watcher?
	// Also, we don't currently have a scenario where we remove a watcher, but we
	// should think through this.
	if cachedWatcher, ok := s.watchers[stackID]; ok {
		// this should never happen

		cachedWatcher.Stop() // make sure we stop listening
		fmt.Println("WARNING: we found a watcher for a tenant that was missing tenantInfo")
	}

	// queue watcher
	s.watchers[stackID] = w

	return ctx
}

// TODO: Finish me. Don't need for demo.
// Watches config maps for changes to tenant configs and removing tenants.
// Context is so we can cancel this blocking call, not a web context
//func (s *serviceImpl) WatchConfigMaps(ctx context.Context) {
//  agg := make(chan string)

//  // collect channels from all watchers
//  var chans []channel
//  for _, c := range s.watchers {
//    chans := append(chans, c)
//  }

//  // all config updates are handled the same. aggregaggregate them into a single channel we can use to
//  // receive
//  for _, ch := range chans {
//    go func(c chan string) {
//      for msg := range c {
//        agg <- msg
//      }
//    }(ch)
//  }

//  select {
//  case msg <- agg:
//    fmt.Println("received ", msg)
//  }
//}

// Builds the kubernetes stackname
func stackName(stackID int64) string {
	return fmt.Sprintf("%d-hackathon-mthg", stackID)
}

// takes INI, initializes db connection and returns it
func initializeDBConnection(config any) *session.SessionDB {
	//~magic~ to connect to db with config
	return nil
}

// Builds tenant info and returns it to enter into cache
func buildTenantInfoFromConfig(stackID int64, config any) *TenantInfo {
	dbCon := initializeDBConnection(config)

	return &TenantInfo{
		StackID:   stackID,
		SessionDB: dbCon,
	}
}
