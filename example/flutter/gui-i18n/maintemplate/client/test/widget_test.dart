// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility that Flutter provides. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:maintemplate/core/i18n/app_localization.dart';

import 'package:maintemplate/main.dart';
import 'package:maintemplate/widgets/widgets/third_translated_widget.dart';

void main() {
  testWidgets('Test locale EN', (WidgetTester tester) async {
    var locale = Locale("en", "US");
    await tester.pumpWidget(Localizations(
      delegates: [
        AppLocalizationsDelegate(locale),
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate
      ],
      locale: locale,
      child: ThirdTranslatedWidget(),
    ));
    await tester.pumpAndSettle();
    final englishText =
        find.text('Third one! Hey google translate, what\'s up?');
    // find twice, original one and translated one
    expect(englishText, findsNWidgets(2));
  });

  testWidgets('Test locale DE', (WidgetTester tester) async {
    var locale = Locale("de");
    await tester.pumpWidget(Localizations(
      delegates: [
        AppLocalizationsDelegate(locale),
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate
      ],
      locale: locale,
      child: ThirdTranslatedWidget(),
    ));
    await tester.pumpAndSettle();

    final englishText =
        find.text('Third one! Hey google translate, what\'s up?');
    expect(englishText, findsOneWidget);

    final germanText = find.text('Dritte! Hey google translate, was ist los?');
    expect(germanText, findsOneWidget);
  });

  testWidgets('Test locale ES', (WidgetTester tester) async {
    var locale = Locale("es");
    await tester.pumpWidget(Localizations(
      delegates: [
        AppLocalizationsDelegate(locale),
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate
      ],
      locale: locale,
      child: ThirdTranslatedWidget(),
    ));
    await tester.pumpAndSettle();

    final englishText =
        find.text('Third one! Hey google translate, what\'s up?');
    expect(englishText, findsOneWidget);

    final spanishText =
        find.text('¡Tercera! Hola traductor de google, ¿qué pasa?');
    expect(spanishText, findsOneWidget);
  });
}
