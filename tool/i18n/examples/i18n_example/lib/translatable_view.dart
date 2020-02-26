

import 'package:flutter/material.dart';
import 'package:i18n_example/i18n/app_localization.dart';

class TranslateAbleView extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title : Text("Traslation Example", style: TextStyle(color : Colors.black), ),
        centerTitle: true,
        backgroundColor: Colors.white,
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Center(child: Column(
          mainAxisSize: MainAxisSize.min,
          children: <Widget>[
            Text(AppLocalizations.of(context).mockText(),),
            SizedBox(height : 50),
            Text("Supported languages -- en, es, fr, ur", style: TextStyle(fontWeight: FontWeight.bold)),
          ],
        )),
      ),
    );
  }
}