import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import "./App.css";
import reportWebVitals from './reportWebVitals';
import {createHashRouter, RouterProvider} from "react-router-dom";
import {NextUIProvider} from "@nextui-org/react";
import Layout from "./layout/layout"
import Search from "./pages/search/search";
import Live from "./pages/live/live";
import {Provider} from 'react-redux';
import store from "./store/store"

const router = createHashRouter([{
    element: <Layout/>, children: [{
        path: "/datapoints/search", element: <Search/>,
    },{
        path: "/", element: <Search/>,
    }, {
        path: "/datapoints/live", element: <Live/>,
    },],
}])


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<Provider store={store}>
    <RouterProvider router={router}>
        <NextUIProvider/>
    </RouterProvider>
</Provider>);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
