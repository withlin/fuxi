package workload

import (
	"context"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
	"io"
	"strconv"
	"time"
)

// Pod doc kubernetes
type Pod struct {
	common.WorkloadsResourceHandler
}

func (p *Pod) Logs(
	namespace, name, container string,
	follow bool, previous bool, timestamps bool,
	sinceSeconds int64,
	sinceTime *time.Time,
	limitBytes int64,
	tailLines int64,
	out io.Writer,
) error {
	req := common.SharedK8sClient.
		ClientV1.
		CoreV1().
		RESTClient().
		Get().
		Namespace(namespace).
		Name(name).
		Resource("pods").
		SubResource("log").
		Param("container", container).
		Param("follow", strconv.FormatBool(follow)).
		Param("previous", strconv.FormatBool(previous)).
		Param("timestamps", strconv.FormatBool(timestamps))

	if sinceSeconds != 0 {
		req.Param("sinceSeconds", strconv.FormatInt(sinceSeconds, 10))
	}
	if sinceTime != nil {
		req.Param("sinceTime", sinceTime.Format(time.RFC3339))
	}
	if limitBytes != 0 {
		req.Param("limitBytes", strconv.FormatInt(limitBytes, 10))
	}
	if tailLines != 0 {
		req.Param("tailLines", strconv.FormatInt(tailLines, 10))
	}
	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return err
	}
	defer readCloser.Close()
	_, err = io.Copy(out, readCloser)

	return err
}

func NewPod() *Pod {
	return &Pod{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourcePod,
	}}
}
