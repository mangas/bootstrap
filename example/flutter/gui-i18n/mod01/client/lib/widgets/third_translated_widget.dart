import 'package:flutter/material.dart';
import 'package:mod01/core/i18n/app_localization.dart';
import 'package:mod01/main.dart';
import 'package:mod01/widgets/translated_text_widget.dart';

class ThirdTranslatedWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return TranslatedTextWidget(
      originalText: ExampleTranslations().testMessageThree(),
      translation: AppLocalizations.of(context).translate("testMessageThree"),
    );
  }
}
