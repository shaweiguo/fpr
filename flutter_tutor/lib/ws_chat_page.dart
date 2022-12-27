import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class WsChatPage extends StatefulWidget {
  const WsChatPage({Key? key}) : super(key: key);

  @override
  State<WsChatPage> createState() => _WsChatPageState();
}

class _WsChatPageState extends State<WsChatPage> {
  final TextEditingController _controller = TextEditingController();
  final _channel = WebSocketChannel.connect(Uri.parse('ws://127.0.0.1:7890'));

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Websocket Chat'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Form(
              child: TextFormField(
                controller: _controller,
                decoration: const InputDecoration(labelText: 'Send a message'),
              ),
            ),
            const SizedBox(height: 24,),
            StreamBuilder(
              stream: _channel.stream,
              builder: (context, snapshot) {
                return Text(snapshot.hasData ? '${snapshot.data}' : '');
              }
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _sendMessage,
        tooltip: 'Send message',
        child: const Icon(Icons.send),
      ),
    );
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      _channel.sink.add(_controller.text);
    }
  }

  @override
  void dispose() {
    _channel.sink.close();
    _controller.dispose();
    super.dispose();
  }
}
