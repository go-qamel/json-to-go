pragma Singleton

import QtQuick 2.11

Item {
    FontLoader {
        id: umRegular
        source: "UbuntuMono-R.ttf"
    }

    FontLoader {
        id: umBold
        source: "UbuntuMono-B.ttf"
    }

    readonly property string bold: umBold.name
    readonly property string regular: umRegular.name
}
