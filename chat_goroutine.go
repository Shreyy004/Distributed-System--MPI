package main

import (
	"fmt"
	"sync"
	"time"
)

// Message struct
type Message struct {
	Sender    string
	Recipient string
	Content   string
	Timestamp time.Time
}

// User struct representing a chat user
type User struct {
	Name       string
	MessageCh  chan Message // Channel for receiving messages
	OfflineMsg []Message    // Store undelivered messages
	mu         sync.Mutex   // Protect OfflineMsg slice
	Online     bool         // Track online status
}

// ChatServer struct
type ChatServer struct {
	Users map[string]*User
	mu    sync.Mutex
}

// Create a new ChatServer
func NewChatServer() *ChatServer {
	return &ChatServer{Users: make(map[string]*User)}
}

// Add a user to the server
func (server *ChatServer) AddUser(name string) {
	server.mu.Lock()
	defer server.mu.Unlock()
	if _, exists := server.Users[name]; !exists {
		server.Users[name] = &User{Name: name, MessageCh: make(chan Message, 10), Online: true}
		fmt.Println("User", name, "added.")
	}
}

// Send a message from one user to another
func (server *ChatServer) SendMessage(sender, recipient, content string) {
	server.mu.Lock()
	user, exists := server.Users[recipient]
	server.mu.Unlock()

	msg := Message{Sender: sender, Recipient: recipient, Content: content, Timestamp: time.Now()}

	if exists {
		user.mu.Lock()
		defer user.mu.Unlock()
		if user.Online {
			// Send message to the active channel
			select {
			case user.MessageCh <- msg:
				fmt.Printf("Message from %s to %s: %s\n", sender, recipient, content)
			default:
				// If channel is full, store the message for later
				user.OfflineMsg = append(user.OfflineMsg, msg)
				fmt.Printf("User %s is busy. Message stored for later delivery.\n", recipient)
			}
		} else {
			// Store message since user is offline
			user.OfflineMsg = append(user.OfflineMsg, msg)
			fmt.Printf("User %s is offline. Message stored for later delivery.\n", recipient)
		}
	} else {
		fmt.Printf("Recipient %s not found. Message lost.\n", recipient)
	}
}

// Process messages for a user (goroutine)
func (user *User) ListenForMessages() {
	for msg := range user.MessageCh {
		fmt.Printf("[%s] %s: %s\n", msg.Timestamp.Format("15:04:05"), msg.Sender, msg.Content)
	}
}

// Simulate a user going offline
func (server *ChatServer) SetUserOffline(name string) {
	server.mu.Lock()
	user, exists := server.Users[name]
	server.mu.Unlock()

	if exists {
		user.mu.Lock()
		user.Online = false
		user.MessageCh = nil // Avoid closing the channel
		user.mu.Unlock()
		fmt.Println(name, "is now offline.")
	}
}

// Simulate a user coming online and receiving stored messages
func (server *ChatServer) SetUserOnline(name string) {
	server.mu.Lock()
	user, exists := server.Users[name]
	server.mu.Unlock()

	if exists {
		user.mu.Lock()
		user.Online = true
		user.MessageCh = make(chan Message, 10) // Reinitialize the channel
		user.mu.Unlock()
		go user.ListenForMessages() // Restart listening
		fmt.Println(name, "is back online.")

		// Deliver stored messages
		server.DeliverStoredMessages(name)
	}
}

// Deliver stored messages
func (server *ChatServer) DeliverStoredMessages(name string) {
	server.mu.Lock()
	user, exists := server.Users[name]
	server.mu.Unlock()

	if exists {
		user.mu.Lock()
		for _, msg := range user.OfflineMsg {
			select {
			case user.MessageCh <- msg:
				fmt.Printf("Delivering stored message to %s: %s\n", name, msg.Content)
			default:
				fmt.Printf("User %s's message queue is full. Cannot deliver all stored messages.\n", name)
				break
			}
		}
		user.OfflineMsg = nil // Clear stored messages
		user.mu.Unlock()
	}
}

func main() {
	server := NewChatServer()

	// Add users
	server.AddUser("Alice")
	server.AddUser("Bob")

	// Start listening for messages (Run in goroutines)
	go server.Users["Alice"].ListenForMessages()
	go server.Users["Bob"].ListenForMessages()

	// Send messages
	server.SendMessage("Alice", "Bob", "Hello, Bob!")
	server.SendMessage("Bob", "Alice", "Hey, Alice!")
	server.SendMessage("Alice", "Bob", "How are you?")
	server.SendMessage("Bob", "Alice", "I'm good. What about you?")

	// Simulate Bob going offline
	time.Sleep(2 * time.Second)
	server.SetUserOffline("Bob")

	// Alice sends messages while Bob is offline
	server.SendMessage("Alice", "Bob", "Are you there?")
	server.SendMessage("Alice", "Bob", "Ping!")

	// Simulate Bob coming back online
	time.Sleep(3 * time.Second)
	server.SetUserOnline("Bob")

	time.Sleep(2 * time.Second) // Allow goroutines to finish
}
