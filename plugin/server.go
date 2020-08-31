package plugin

import (
	"log"
	"net"
)

type Server struct {
	Net      string
	Addr     string
	listener net.Listener
}

func (s *Server) Listen() error {
	var err error
	s.listener, err = net.Listen(s.Net, s.Addr)
	if err != nil {
		return err
	}
	go s.accept()
	return nil
}

func (s *Server) Close() error {
	if s.listener != nil {
		return nil
	}
	err := s.listener.Close()
	if err != nil {
		return err
	}
	s.listener = nil
	return nil
}

func (s *Server) accept() {
	for s.listener != nil {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
			//TODO 判断监听异常应该退出
		}
		go s.receive(conn)
	}
}

func (s *Server) receive(conn net.Conn) {
	//TODO 接收Key，并校验

	buf := make([]byte, 1024)
	for {
		n, e := conn.Read(buf)
		//TODO 解析消息，并处理

	}
}