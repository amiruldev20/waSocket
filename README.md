# waSocket
[![Go Reference](https://pkg.go.dev/badge/github.com/amiruldev20/waSocket.svg)](https://pkg.go.dev/github.com/amiruldev20/waSocket)

waSocket is a Go library for the WhatsApp web multidevice API.

## Discussion
Matrix room: [#waSocket:maunium.net](https://matrix.to/#/#waSocket:maunium.net)

For questions about the WhatsApp protocol (like how to send a specific type of
message), you can also use the [WhatsApp protocol Q&A] section on GitHub
discussions.

[WhatsApp protocol Q&A]: https://github.com/tulir/waSocket/discussions/categories/whatsapp-protocol-q-a

## Usage
The [godoc](https://pkg.go.dev/github.com/amiruldev20/waSocket) includes docs for all methods and event types.
There's also a [simple example](https://godocs.io/github.com/amiruldev20/waSocket#example-package) at the top.

Also see [mdtest](./mdtest) for a CLI tool you can easily try out waSocket with.

## Features
Most core features are already present:

* Sending messages to private chats and groups (both text and media)
* Receiving all messages
* Managing groups and receiving group change events
* Joining via invite messages, using and creating invite links
* Sending and receiving typing notifications
* Sending and receiving delivery and read receipts
* Reading and writing app state (contact list, chat pin/mute status, etc)
* Sending and handling retry receipts if message decryption fails
* Sending status messages (experimental, may not work for large contact lists)

Things that are not yet implemented:

* Sending broadcast list messages (this is not supported on WhatsApp web either)
* Calls
