pragma Singleton

import QtQuick 2.11

Item {
    FontLoader {
        id: faSolid
        source: "Font Awesome 5 Free-Solid-900.otf"
    }

    FontLoader {
        id: faRegular
        source: "Font Awesome 5 Free-Regular-400.otf"
    }

    FontLoader {
        id: faBrands
        source: "Font Awesome 5 Brands-Regular-400.otf"
    }

    readonly property string solid: faSolid.name
    readonly property string brands: faBrands.name
    readonly property string regular: faRegular.name
}