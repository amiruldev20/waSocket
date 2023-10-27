package types

import "time"

type BasicCallMeta struct {
	From        JID
	Timestamp   time.Time
	CallCreator JID
	CallID      string
}

type CallRemoteMeta struct {
	RemotePlatform string // The platform of the caller's WhatsApp client
	RemoteVersion  string // Version of the caller's WhatsApp client
}
