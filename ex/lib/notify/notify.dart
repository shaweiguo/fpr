import 'package:flutter/material.dart';
import 'package:local_notifier/local_notifier.dart';
// import 'package:bot_toast/bot_toast.dart';

class Notify extends StatefulWidget {
  const Notify({Key? key}) : super(key: key);

  @override
  State<Notify> createState() => _NotifyState();
}

class _NotifyState extends State<Notify> {
  void _notify() {
    final notification = LocalNotification(
      identifier: "00001",
      title: "古诗词鉴赏",
      subtitle: "桃夭 - 佚名〔先秦〕",
      body: '桃之夭夭，灼灼其华。之子于归，宜其室家。\n桃之夭夭，有蕡其实。之子于归，宜其家室。\n桃之夭夭，其叶蓁蓁。之子于归，宜其家人。',
      // 用来设置是否静音
      silent: true,
      actions: [
        LocalNotificationAction(text: '学废了'),
        LocalNotificationAction(text: '没学废'),
      ],
    );

    // notification.onShow = () => BotToast.showText(text: "显示一条通知。");
    // notification.onClose = (reason) => BotToast.showText(text: "由于$reason，通知关闭！");
    // notification.onClick = () => BotToast.showText(text: "通知被点击。");
    // notification.onClickAction = (index) => BotToast.showText(text: "选择${notification.actions?[index].text}");
    notification.onShow = () => print("显示一条通知。");
    notification.onClose = (reason) => print("由于$reason，通知关闭！");
    notification.onClick = () => print("通知被点击。");
    notification.onClickAction = (index) => print("选择${notification.actions?[index].text}");
    
    notification.show();
  }
  
  @override
  Widget build(BuildContext context) {
    return ElevatedButton(onPressed: _notify, child: const Text("发送通知"));
  }
}
