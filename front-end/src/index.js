import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import "./App.css";
import reportWebVitals from './reportWebVitals';
import {RouterProvider, createBrowserRouter} from "react-router-dom";
import {NextUIProvider} from "@nextui-org/react";
import Layout from "./layout/layout"

// Lazy load
const Search = React.lazy(() => import("./pages/search/search"))

const router = createBrowserRouter([
    {
        element: <Layout/>,
        children:[
            {
                path: "/datapoints/search",
                element: <Search />,
            },
        ],
    }
])
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <RouterProvider router={router}>
        <NextUIProvider/>
    </RouterProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
