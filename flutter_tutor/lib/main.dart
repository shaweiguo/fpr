import 'package:flutter/material.dart';

import 'favorite_widget.dart';
import 'ws_chat_page.dart';
import 'video_player_page.dart';
import 'file_handle_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Layout Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      routes: {
        'chat': (context) => const WsChatPage(),
        'play': (context) => const VideoPlayerPage(),
        'file': (context) => FileHandlePage(storage: MyFileStorage()),
      },
      home: const MyHomePage(title: 'Flutter Layout Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }

  Widget _buildTitleSection() => Container(
      padding: const EdgeInsets.all(32),
      child: Row(
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  padding: const EdgeInsets.only(bottom: 8),
                  child: const Text(
                    '凤凰桥',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                Text(
                  '中国, 宁夏, 银川',
                  style: TextStyle(
                    color: Colors.grey[500],
                  ),
                ),
              ],
            ),
          ),
          const FavoriteWidget(),
          Text(
            '$_counter',
            // style: Theme.of(context).textTheme.headline6,
          ),
        ],
      ),
    );

  Column _buildButtonColumn(Color color, IconData icon, String label) => Column(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Icon(icon, color: color),
        Container(
          margin: const EdgeInsets.only(top: 8),
          child: Text(
            label,
            style: TextStyle(
              fontSize: 12,
              fontWeight: FontWeight.w400,
              color: color,
            ),
          ),
        ),
      ],
    );
  
  Widget _buildButton(label, route, Icon icon) => ElevatedButton.icon(
    onPressed: () => Navigator.pushNamed(context, route),
    icon: icon,
    label: Text(label),
  );

  Widget _buildButtonSection(Color color) => Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        _buildButtonColumn(color, Icons.call, 'CALL'),
        _buildButtonColumn(color, Icons.near_me, 'ROUTE'),
        // _buildButtonColumn(color, Icons.share, 'SHARE'),
        _buildButton("聊天", "chat", const Icon(Icons.chat)),
        _buildButton("视频", "play", const Icon(Icons.video_file)),
        _buildButton("文件", "file", const Icon(Icons.file_open)),
      ],
    );

  Widget _buildTextSection() => const Padding(
      padding: EdgeInsets.all(32),
      child: Text(
        '凤凰桥是凤城银川又一座横跨典农河的重要桥梁，是典农河上靓丽的地标景观。2018年建成后有效缓解了正源街和民族街南北向的交通压力。凤凰桥采用一跨钢拱桥，行车道布置两片内倾主拱，人行道布置两片外倾主拱，内外拱通过水平拉索进行受力平衡，形成极具张力的展翼造型，达到结构和艺术的和谐一致，突显了凤凰文化内涵。',
        softWrap: true,
      ),
    );

  @override
  Widget build(BuildContext context) {
    Color color = Theme.of(context).primaryColor;

    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: ListView(
          // mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Image.asset(
              'assets/images/yc_fhq1.jpg',
              width: 600,
              height: 240,
              fit: BoxFit.cover,
            ),
            _buildTitleSection(),
            _buildButtonSection(color),
            _buildTextSection(),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
