import QtQuick 2.12
import QtQuick.Controls 2.12
import "../themes" as T
import "../fonts/UbuntuMono" as UM

TextArea {
    id: root

    padding: 16
    color: T.Dracula.font
    selectionColor: T.Dracula.selectionBg
    selectedTextColor: color
    selectByMouse: true
    selectByKeyboard: true
    font.pointSize: 12
    font.family: UM.Fonts.regular
    tabStopDistance: 24

    background: Rectangle {
        color: T.Dracula.contentBg
        anchors.fill: parent
    }
}