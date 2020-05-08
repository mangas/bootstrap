import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:mod01/core/i18n/app_localization.dart';
import 'package:mod01/core/i18n/translations.dart';
import 'package:mod01/widgets/first_translated_widget.dart';
import 'package:mod01/widgets/second_translated_widget.dart';
import 'package:mod01/widgets/third_translated_widget.dart';
import 'package:mod01/widgets/translated_text_widget.dart';

void main() => runApp(App());

class App extends StatefulWidget {
  @override
  _AppState createState() => _AppState();
}

class ExampleTranslations extends Translations {}

class _AppState extends State<App> {
  var _locale = Locale("en", "US");

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        debugShowCheckedModeBanner: false,
        localizationsDelegates: [
          AppLocalizationsDelegate(_locale),
          GlobalMaterialLocalizations.delegate,
          GlobalWidgetsLocalizations.delegate,
        ],
        locale: _locale,
        supportedLocales: AppLocalizations.supportedLanguages.keys
            .map((e) => Locale(e))
            .toList(),
        home: Builder(builder: (context) {
          return Scaffold(
            appBar: AppBar(
              title: Text("Mod01 i18n Example"),
            ),
            body: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(
                  child: SingleChildScrollView(
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: <Widget>[
                    FirstTranslatedWidget(),
                    SizedBox(
                      height: 32,
                    ),
                    SecondTranslatedWidget(),
                    SizedBox(
                      height: 32,
                    ),
                    ThirdTranslatedWidget(),
                    SizedBox(height: 50),
                    Text(
                        "Supported languages : ${AppLocalizations.supportedLanguages.entries.map((e) => e.key).toList()}",
                        style: TextStyle(fontWeight: FontWeight.bold)),
                    SizedBox(height: 50),
                    Wrap(
                      alignment: WrapAlignment.spaceEvenly,
                      runAlignment: WrapAlignment.spaceEvenly,
                      children: <Widget>[
                        ...AppLocalizations.supportedLanguages.keys
                            .map((e) => RaisedButton(
                                  child: Text(e),
                                  onPressed: () {
                                    setState(() {
                                      _locale = Locale(e, e);
                                    });
                                  },
                                ))
                            .toList(),
                      ],
                    )
                  ],
                ),
              )),
            ),
          );
        }));
  }
}
