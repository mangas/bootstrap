import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:i18n_example/core/i18n/app_localization.dart';

void main() => runApp(App());

class App extends StatefulWidget {
  @override
  _AppState createState() => _AppState();
}

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
        supportedLocales: AppLocalizations.getLocales(),
        home: Builder(builder: (context) {
          return Scaffold(
            body: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(
                  child: Column(
                mainAxisSize: MainAxisSize.min,
                children: <Widget>[
                  Text(
                    AppLocalizations.of(context).testMessage(),
                  ),
                  SizedBox(height: 50),
                  Text(
                      "Supported languages : ${AppLocalizations.supportedLanguages.entries.map((e) => e.key).toList()}",
                      style: TextStyle(fontWeight: FontWeight.bold)),
                  SizedBox(height: 50),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: <Widget>[
                      RaisedButton(
                        child: Text("EN"),
                        onPressed: () {
                          setState(
                            () {
                              _locale = Locale("en", "en");
                            },
                          );
                        },
                      ),
                      RaisedButton(
                        child: Text("ES"),
                        onPressed: () {
                          setState(
                            () {
                              _locale = Locale("es", "es");
                            },
                          );
                        },
                      ),
                      RaisedButton(
                        child: Text("FR"),
                        onPressed: () {
                          setState(() {
                            _locale = Locale("fr", "fr");
                          });
                        },
                      ),
                      RaisedButton(
                        child: Text("DE"),
                        onPressed: () {
                          setState(() {
                            _locale = Locale("de", "de");
                          });
                        },
                      ),
                      RaisedButton(
                        child: Text("UR"),
                        onPressed: () {
                          setState(
                            () {
                              _locale = Locale("ur", "ur");
                            },
                          );
                        },
                      ),
                    ],
                  )
                ],
              )),
            ),
          );
        }));
  }
}
