# net

### 监听方法

Listener类型面向流的协议的通用监听器，多个g可以同时调用Listener上的方法Accept、Close、Addr

- ```go
  func FileListener(f *os.File) (ln Listener, err error)
  ```

- Listen、ListenIP

### 连接方法

Dial、DialIP

### 数据类型

##### IP地址

- ParseCIDR解析字符串

### 连接类型

- Conn：通用的面向流的网络连接，TCPConn实现该接口
- PacketConn：通用面向数据包的网络连接，UDPConn和IPConn实现了该接口

### 连接读写 Conn

#### 通用方法

读：ReadFrom(返回读取的字节数)、ReadFromIP(返回一个IPAddr)

写：WriteTo(参数类型为Addr)、WriteToIP(参数类型为*IPAddr)

Set(Read/Write)Buffer设置读取缓冲区大小

#### 面向流的连接方法

Read、Write

Set(Read/Write)Deadline、

#### IPConn

- ```go
  func (c *IPConn) File() (f *os.File, err error) 获得这个网络连接的文件描述符
  ```

- ```go
  func (c *IPConn) SyscallConn() (syscall.RawConn, error) 获得这个网络连接的系统层原生连接
  ```

- ```go
  func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
  ```

- ```go
  func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
  ```

- 





#### UDPConn

- 