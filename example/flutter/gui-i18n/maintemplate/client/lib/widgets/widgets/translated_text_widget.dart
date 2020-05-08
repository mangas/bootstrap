import 'package:flutter/material.dart';

class TranslatedTextWidget extends StatelessWidget {
  final String originalText;
  final String translation;

  const TranslatedTextWidget(
      {Key key, @required this.originalText, @required this.translation})
      : assert(originalText != null),
        assert(translation != null),
        super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Text("Not translated original String:"),
        Text(originalText,
            style: TextStyle(
              color: Colors.red,
            )),
        SizedBox(
          height: 8,
        ),
        Text("Translated String:"),
        Text(translation,
            style: TextStyle(
              color: Colors.green,
            ))
      ],
    );
  }
}
