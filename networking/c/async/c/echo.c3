import std::io;
import std::net::tcp;

fn void main() {
  TcpServerSocket server = tcp::listen("localhost", 6969, 69, REUSEADDR)!!; // 69 queue size
  
  while (true) {
    TcpSocket client = tcp::accept(&server)!!;
    io::printfn("OK");
    client.close();
  }

  io::printfn("Hello, World");
}
