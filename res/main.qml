import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12
import "themes" as T
import "components"
import "fonts/FontAwesome" as FA
import "fonts/SourceSansPro" as SSP
import BackEnd 1.0

Rectangle {
    color: T.Dracula.bg

    DocHandler {
        id: docHandler

        function convertJSON() {
            if (txtJson.text === "") txtGo.text = "";
            else docHandler.convert(txtJson.text, chkInline.checked)
        }

        onConverted: (val) => {
            txtGo.text = val;
        }
        onError: (err) => {
            txtGo.text = `<font color="#FF6666">${err}</font>`;
        }
    }

    ColumnLayout {
        spacing: 0
        anchors.fill: parent

        Rectangle {
            color: T.Dracula.bg
            height: headerContent.height
			Layout.fillWidth: true

            RowLayout {
                id: headerContent
                spacing: 0
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.right: parent.right

                Text {
                    text: "JSON"
                    padding: 8
                    color: T.Dracula.font
                    font.family: SSP.Fonts.semiBold
                    font.pointSize: 15
                    Layout.fillWidth: true
                    Layout.preferredWidth: 1
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                }

                Text {
                    text: FA.Icons.faArrowRight
                    color: T.Dracula.font
                    font.weight: Font.Bold
                    font.family: FA.Fonts.solid
                    font.pointSize: 12
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                }

                Text {
                    text: "Go"
                    padding: 8
                    color: T.Dracula.font
                    font.family: SSP.Fonts.semiBold
                    font.pointSize: 15
                    Layout.fillWidth: true
                    Layout.preferredWidth: 1
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                }
            }
        }

        Rectangle {
            height: 1
            color: T.Dracula.border
            Layout.fillWidth: true
        }

        Rectangle {
            color: "transparent"
            Layout.fillWidth: true
            Layout.fillHeight: true

            RowLayout {
                spacing: 0
                anchors.fill: parent

                ScrollView {
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.preferredWidth: 1

                    CustomTextArea {
                        id: txtJson
                        placeholderText: "Paste JSON here"
                        onTextChanged: docHandler.convertJSON()
                    }
                }

                Rectangle {
                    width: 1
                    color: T.Dracula.border
                    Layout.fillHeight: true
                }

                ScrollView {
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.preferredWidth: 1

                    CustomTextArea {
                        id: txtGo
                        readOnly: true
                        placeholderText: "Go will appear here"
                        textFormat: TextEdit.RichText
                    }
                }
            }
        }

        Rectangle {
            height: 1
            color: T.Dracula.border
            Layout.fillWidth: true
        }

        Rectangle {
            color: T.Dracula.bg
            height: footerContent.height
			Layout.fillWidth: true

            RowLayout {
                id: footerContent
                spacing: 0
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.right: parent.right

                CustomCheckBox {
                    id: chkInline

                    padding: 8
                    spacing: 8
                    text: "Inline struct"
                    height: parent.height
                    font.family: SSP.Fonts.regular
                    font.pointSize: 10
                    onCheckedChanged: docHandler.convertJSON()
                }

                Text {
                    padding: 8
                    color: T.Dracula.font
                    font.family: SSP.Fonts.regular
                    font.pointSize: 10
                    Layout.fillWidth: true
                    Layout.preferredWidth: 1
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignRight
                    text: "Original web app © 2015 Matt Holt, ported to Go + QML by RadhiFadlillah"
                }
            }
        }
    }
}
