import 'package:flutter/material.dart';
import 'package:maintemplate/core/i18n/app_localization.dart';
import 'package:maintemplate/main.dart';
import 'package:maintemplate/widgets/widgets/translated_text_widget.dart';

class FirstTranslatedWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return TranslatedTextWidget(
      originalText: ExampleTranslations().testMessage(),
      translation: AppLocalizations.of(context).testMessage(),
    );
  }
}
