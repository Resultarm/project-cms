"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(() => {
var exports = {};
exports.id = "pages/admin";
exports.ids = ["pages/admin"];
exports.modules = {

/***/ "./pages/admin.js":
/*!************************!*\
  !*** ./pages/admin.js ***!
  \************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react */ \"react\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var _services_api__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/api */ \"./services/api.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"react/jsx-dev-runtime\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__);\nvar _jsxFileName = \"/Users/franciscosantos/go/src/Works/rmcsm/linktreeResultar/pages/admin.js\";\n\n\n\n\nfunction Admin() {\n  const {\n    0: username,\n    1: setUsername\n  } = (0,react__WEBPACK_IMPORTED_MODULE_0__.useState)(\"\");\n  const {\n    0: password,\n    1: setPassword\n  } = (0,react__WEBPACK_IMPORTED_MODULE_0__.useState)(\"\");\n\n  const submit = e => {\n    e.preventDefault();\n\n    if (username == \"\" || password == \"\") {\n      return;\n    }\n\n    let data = {\n      username: username,\n      password: password\n    };\n    _services_api__WEBPACK_IMPORTED_MODULE_1__.default.post(\"/login\", data, {\n      headers: {\n        \"Content-Type\": \"application/json\"\n      }\n    }).then(r => {\n      localStorage.setItem(\"token\", r.data.token);\n      window.location.href = \"/dashboard\";\n    });\n  };\n\n  return /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__.jsxDEV)(\"div\", {\n    children: /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__.jsxDEV)(\"form\", {\n      onSubmit: submit,\n      children: [/*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__.jsxDEV)(\"input\", {\n        type: \"text\",\n        onChange: e => setUsername(e.target.value)\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 25,\n        columnNumber: 9\n      }, this), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__.jsxDEV)(\"input\", {\n        type: \"password\",\n        onChange: e => setPassword(e.target.value)\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 26,\n        columnNumber: 9\n      }, this), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_2__.jsxDEV)(\"input\", {\n        type: \"submit\"\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 30,\n        columnNumber: 9\n      }, this)]\n    }, void 0, true, {\n      fileName: _jsxFileName,\n      lineNumber: 24,\n      columnNumber: 7\n    }, this)\n  }, void 0, false, {\n    fileName: _jsxFileName,\n    lineNumber: 23,\n    columnNumber: 5\n  }, this);\n}\n\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (Admin);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9hZG1pbi5qcy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7O0FBQUE7QUFDQTs7O0FBRUEsU0FBU0UsS0FBVCxHQUFpQjtBQUNmLFFBQU07QUFBQSxPQUFDQyxRQUFEO0FBQUEsT0FBV0M7QUFBWCxNQUEwQkosK0NBQVEsQ0FBQyxFQUFELENBQXhDO0FBQ0EsUUFBTTtBQUFBLE9BQUNLLFFBQUQ7QUFBQSxPQUFXQztBQUFYLE1BQTBCTiwrQ0FBUSxDQUFDLEVBQUQsQ0FBeEM7O0FBRUEsUUFBTU8sTUFBTSxHQUFJQyxDQUFELElBQU87QUFDcEJBLElBQUFBLENBQUMsQ0FBQ0MsY0FBRjs7QUFDQSxRQUFJTixRQUFRLElBQUksRUFBWixJQUFrQkUsUUFBUSxJQUFJLEVBQWxDLEVBQXNDO0FBQ3BDO0FBQ0Q7O0FBRUQsUUFBSUssSUFBSSxHQUFHO0FBQUVQLE1BQUFBLFFBQVEsRUFBRUEsUUFBWjtBQUFzQkUsTUFBQUEsUUFBUSxFQUFFQTtBQUFoQyxLQUFYO0FBQ0FKLElBQUFBLHVEQUFBLENBQ1EsUUFEUixFQUNrQlMsSUFEbEIsRUFDd0I7QUFBRUUsTUFBQUEsT0FBTyxFQUFFO0FBQUUsd0JBQWdCO0FBQWxCO0FBQVgsS0FEeEIsRUFFR0MsSUFGSCxDQUVTQyxDQUFELElBQU87QUFDWEMsTUFBQUEsWUFBWSxDQUFDQyxPQUFiLENBQXFCLE9BQXJCLEVBQThCRixDQUFDLENBQUNKLElBQUYsQ0FBT08sS0FBckM7QUFDQUMsTUFBQUEsTUFBTSxDQUFDQyxRQUFQLENBQWdCQyxJQUFoQixHQUF1QixZQUF2QjtBQUNELEtBTEg7QUFNRCxHQWJEOztBQWNBLHNCQUNFO0FBQUEsMkJBQ0U7QUFBTSxjQUFRLEVBQUViLE1BQWhCO0FBQUEsOEJBQ0U7QUFBTyxZQUFJLEVBQUUsTUFBYjtBQUFxQixnQkFBUSxFQUFHQyxDQUFELElBQU9KLFdBQVcsQ0FBQ0ksQ0FBQyxDQUFDYSxNQUFGLENBQVNDLEtBQVY7QUFBakQ7QUFBQTtBQUFBO0FBQUE7QUFBQSxjQURGLGVBRUU7QUFDRSxZQUFJLEVBQUUsVUFEUjtBQUVFLGdCQUFRLEVBQUdkLENBQUQsSUFBT0YsV0FBVyxDQUFDRSxDQUFDLENBQUNhLE1BQUYsQ0FBU0MsS0FBVjtBQUY5QjtBQUFBO0FBQUE7QUFBQTtBQUFBLGNBRkYsZUFNRTtBQUFPLFlBQUksRUFBRTtBQUFiO0FBQUE7QUFBQTtBQUFBO0FBQUEsY0FORjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFERjtBQUFBO0FBQUE7QUFBQTtBQUFBLFVBREY7QUFZRDs7QUFFRCxpRUFBZXBCLEtBQWYiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9saW5rdHJlZS1jbG9uZS8uL3BhZ2VzL2FkbWluLmpzPzIxM2UiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgdXNlU3RhdGUgfSBmcm9tIFwicmVhY3RcIjtcbmltcG9ydCBhcGkgZnJvbSBcIi4uL3NlcnZpY2VzL2FwaVwiO1xuXG5mdW5jdGlvbiBBZG1pbigpIHtcbiAgY29uc3QgW3VzZXJuYW1lLCBzZXRVc2VybmFtZV0gPSB1c2VTdGF0ZShcIlwiKTtcbiAgY29uc3QgW3Bhc3N3b3JkLCBzZXRQYXNzd29yZF0gPSB1c2VTdGF0ZShcIlwiKTtcblxuICBjb25zdCBzdWJtaXQgPSAoZSkgPT4ge1xuICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICBpZiAodXNlcm5hbWUgPT0gXCJcIiB8fCBwYXNzd29yZCA9PSBcIlwiKSB7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgbGV0IGRhdGEgPSB7IHVzZXJuYW1lOiB1c2VybmFtZSwgcGFzc3dvcmQ6IHBhc3N3b3JkIH07XG4gICAgYXBpXG4gICAgICAucG9zdChcIi9sb2dpblwiLCBkYXRhLCB7IGhlYWRlcnM6IHsgXCJDb250ZW50LVR5cGVcIjogXCJhcHBsaWNhdGlvbi9qc29uXCIgfSB9KVxuICAgICAgLnRoZW4oKHIpID0+IHtcbiAgICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oXCJ0b2tlblwiLCByLmRhdGEudG9rZW4pO1xuICAgICAgICB3aW5kb3cubG9jYXRpb24uaHJlZiA9IFwiL2Rhc2hib2FyZFwiO1xuICAgICAgfSk7XG4gIH07XG4gIHJldHVybiAoXG4gICAgPGRpdj5cbiAgICAgIDxmb3JtIG9uU3VibWl0PXtzdWJtaXR9PlxuICAgICAgICA8aW5wdXQgdHlwZT17XCJ0ZXh0XCJ9IG9uQ2hhbmdlPXsoZSkgPT4gc2V0VXNlcm5hbWUoZS50YXJnZXQudmFsdWUpfSAvPlxuICAgICAgICA8aW5wdXRcbiAgICAgICAgICB0eXBlPXtcInBhc3N3b3JkXCJ9XG4gICAgICAgICAgb25DaGFuZ2U9eyhlKSA9PiBzZXRQYXNzd29yZChlLnRhcmdldC52YWx1ZSl9XG4gICAgICAgIC8+XG4gICAgICAgIDxpbnB1dCB0eXBlPXtcInN1Ym1pdFwifSAvPlxuICAgICAgPC9mb3JtPlxuICAgIDwvZGl2PlxuICApO1xufVxuXG5leHBvcnQgZGVmYXVsdCBBZG1pbjtcbiJdLCJuYW1lcyI6WyJ1c2VTdGF0ZSIsImFwaSIsIkFkbWluIiwidXNlcm5hbWUiLCJzZXRVc2VybmFtZSIsInBhc3N3b3JkIiwic2V0UGFzc3dvcmQiLCJzdWJtaXQiLCJlIiwicHJldmVudERlZmF1bHQiLCJkYXRhIiwicG9zdCIsImhlYWRlcnMiLCJ0aGVuIiwiciIsImxvY2FsU3RvcmFnZSIsInNldEl0ZW0iLCJ0b2tlbiIsIndpbmRvdyIsImxvY2F0aW9uIiwiaHJlZiIsInRhcmdldCIsInZhbHVlIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./pages/admin.js\n");

/***/ }),

/***/ "./services/api.js":
/*!*************************!*\
  !*** ./services/api.js ***!
  \*************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ \"axios\");\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_0__);\n\nconst api = axios__WEBPACK_IMPORTED_MODULE_0___default().create({\n  baseURL: \"http://localhost:8000\"\n});\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (api);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zZXJ2aWNlcy9hcGkuanMuanMiLCJtYXBwaW5ncyI6Ijs7Ozs7O0FBQUE7QUFFQSxNQUFNQyxHQUFHLEdBQUdELG1EQUFBLENBQWE7QUFDdkJHLEVBQUFBLE9BQU8sRUFBRTtBQURjLENBQWIsQ0FBWjtBQUlBLGlFQUFlRixHQUFmIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vbGlua3RyZWUtY2xvbmUvLi9zZXJ2aWNlcy9hcGkuanM/ZDhjNCJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgYXhpb3MgZnJvbSBcImF4aW9zXCI7XG5cbmNvbnN0IGFwaSA9IGF4aW9zLmNyZWF0ZSh7XG4gIGJhc2VVUkw6IFwiaHR0cDovL2xvY2FsaG9zdDo4MDAwXCIsXG59KTtcblxuZXhwb3J0IGRlZmF1bHQgYXBpO1xuIl0sIm5hbWVzIjpbImF4aW9zIiwiYXBpIiwiY3JlYXRlIiwiYmFzZVVSTCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./services/api.js\n");

/***/ }),

/***/ "axios":
/*!************************!*\
  !*** external "axios" ***!
  \************************/
/***/ ((module) => {

module.exports = require("axios");

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "react" ***!
  \************************/
/***/ ((module) => {

module.exports = require("react");

/***/ }),

/***/ "react/jsx-dev-runtime":
/*!****************************************!*\
  !*** external "react/jsx-dev-runtime" ***!
  \****************************************/
/***/ ((module) => {

module.exports = require("react/jsx-dev-runtime");

/***/ })

};
;

// load runtime
var __webpack_require__ = require("../webpack-runtime.js");
__webpack_require__.C(exports);
var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
var __webpack_exports__ = (__webpack_exec__("./pages/admin.js"));
module.exports = __webpack_exports__;

})();