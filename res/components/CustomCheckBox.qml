import QtQuick 2.12
import QtQuick.Controls 2.12
import "../themes" as T
import "../fonts/FontAwesome" as FA
import "../fonts/SourceSansPro" as SSP

CheckBox {
    id: root

    MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        enabled: false
    }

    indicator: Rectangle {
        implicitWidth: implicitHeight
        implicitHeight: root.height - (root.topPadding * 2)
        x: root.leftPadding
        y: root.height / 2 - height / 2
        color: "transparent"
        border.width: 1
        border.color: T.Dracula.border

        Text {
            color: T.Dracula.border
            text: FA.Icons.faTimes
            anchors.fill: parent
            font.pointSize: 11
            font.weight: Font.Bold
            font.family: FA.Fonts.solid
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            visible: root.checked
        }
    }

    contentItem: Text {
        text: root.text
        font: root.font
        opacity: enabled ? 1.0 : 0.3
        color: T.Dracula.font
        verticalAlignment: Text.AlignVCenter
        leftPadding: root.indicator.width + root.spacing
    }
}