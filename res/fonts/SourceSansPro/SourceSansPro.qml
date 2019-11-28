pragma Singleton

import QtQuick 2.11

Item {
    FontLoader {
        id: sspRegular
        source: "Source Sans Pro-Regular.ttf"
    }

    FontLoader {
        id: sspSemiBold
        source: "Source Sans Pro-600.ttf"
    }

    FontLoader {
        id: sspBold
        source: "Source Sans Pro-700.ttf"
    }

    readonly property string bold: sspBold.name
    readonly property string regular: sspRegular.name
    readonly property string semiBold: sspSemiBold.name
}
