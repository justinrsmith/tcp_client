package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net"
    "os"
)

var (
    ip string
    port string
)

func init() {
    flag.StringVar(&ip, "ip", "127.0.0.1", "IP Address to assign the server")
    flag.StringVar(&port, "port", "3333", "Port to assign to the server")
    flag.Parse()
}

func main() {
    for {
        // Reconnect on each pass through since server closes connection
        conn, err := net.Dial("tcp", ip+":"+port)
        if err != nil {
            // Log the error (Print) and exit program
            log.Fatal(err)
        }
        // Read user input
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Provide a message: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        // Send string to server
        conn.Write([]byte(input))
        // Read response from server
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        fmt.Print(message)
    }
}

