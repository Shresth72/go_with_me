import std::io;
import std::net::tcp;

fn void main() {
  TcpServerSocket server = tcp::listen("localhost", 6969, 69, REUSEADDR)!!; // 69 queue size
  
  char[1024] buffer;

  while (true) {
    TcpSocket client = tcp::accept(&server)!!;

    usz n = client.read(&buffer)!!;
    // io::printfn("Client sent: %s", (String)buffer[0:n]);
    
    client.write(buffer[0:n]);
      
    client.close();
  }

  
}
