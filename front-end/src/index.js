import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import "./App.css";
import reportWebVitals from './reportWebVitals';
import {createHashRouter, RouterProvider} from "react-router-dom";
import {NextUIProvider} from "@nextui-org/react";
import LayoutApp from "./layout/app/layout"
import LayoutLogin from "./layout/login/layout"
import Search from "./pages/search/search";
import Live from "./pages/live/live";
import {Provider} from 'react-redux';
import store from "./store/store"
import Login from "./pages/login/login";
import Devices from "./pages/devices/devices";
import AccessManagement from "./pages/accessManagement/accessManagement";
import Main from "./pages/main/main";

const router = createHashRouter([
    {
        element: <LayoutApp/>,
        children: [
            {
                path: "/datapoints/search", element: <Search/>,
            },
            {
                path: "/", element: <Main/>,
            },
            {
                path: "/datapoints/live", element: <Live/>,
            },
            {
                path: "/devices", element: <Devices/>,
            },
            {
                path: "/access-management", element: <AccessManagement/>,
            }
        ],
    },
    {
        element: <LayoutLogin/>,
        children: [
            {
                path: "/login", element: <Login/>,
            },
        ],
    }
])


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <Provider store={store}>
        <RouterProvider router={router}>
            <NextUIProvider/>
        </RouterProvider>
    </Provider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
