// 2023 © Whatsmeow
// Redeveloped by Amirul Dev

package waSocket

import (
	"context"

	waBinary "github.com/amiruldev20/waSocket/binary"
	"github.com/amiruldev20/waSocket/types"
)

type DangerousInternalClient struct {
	c *Client
}

// DangerousInternals allows access to some unexported methods in Client.
//
// Deprecated: dangerous
func (cli *Client) DangerousInternals() *DangerousInternalClient {
	return &DangerousInternalClient{cli}
}

type DangerousInfoQuery = infoQuery
type DangerousInfoQueryType = infoQueryType

func (int *DangerousInternalClient) SendIQ(query DangerousInfoQuery) (*waBinary.Node, error) {
	return int.c.sendIQ(query)
}

func (int *DangerousInternalClient) SendIQAsync(query DangerousInfoQuery) (<-chan *waBinary.Node, error) {
	return int.c.sendIQAsync(query)
}

func (int *DangerousInternalClient) SendNode(node waBinary.Node) error {
	return int.c.sendNode(node)
}

func (int *DangerousInternalClient) WaitResponse(reqID string) chan *waBinary.Node {
	return int.c.waitResponse(reqID)
}

func (int *DangerousInternalClient) CancelResponse(reqID string, ch chan *waBinary.Node) {
	int.c.cancelResponse(reqID, ch)
}

func (int *DangerousInternalClient) QueryMediaConn() (*MediaConn, error) {
	return int.c.queryMediaConn()
}

func (int *DangerousInternalClient) RefreshMediaConn(force bool) (*MediaConn, error) {
	return int.c.refreshMediaConn(force)
}

func (int *DangerousInternalClient) GetServerPreKeyCount() (int, error) {
	return int.c.getServerPreKeyCount()
}

func (int *DangerousInternalClient) RequestAppStateKeys(ctx context.Context, keyIDs [][]byte) {
	int.c.requestAppStateKeys(ctx, keyIDs)
}

func (int *DangerousInternalClient) SendRetryReceipt(node *waBinary.Node, info *types.MessageInfo, forceIncludeIdentity bool) {
	int.c.sendRetryReceipt(node, info, forceIncludeIdentity)
}
