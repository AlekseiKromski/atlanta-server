import {createSlice} from '@reduxjs/toolkit';
import axios from "axios"
import Cookies from 'js-cookie';

const applicationSlice = createSlice({
        name: 'axios',
        initialState: {
            axios: function () {
                let instance = axios.create({
                    baseURL: process.env.REACT_APP_AXIOS_BASE_URL,
                    timeout: 1000,
                    headers: {},
                    withCredentials: true
                })

                instance.interceptors.response.use(
                    (response) => {
                        return response;
                    },
                    (error) => {
                        if (error.response && error.response.status === 401) {
                            // Use router.push() to navigate to the login screen
                            document.location = "/#login"
                            // Throw an exception to stop further execution
                            return Promise.reject(error.response.data.message);
                        }
                        // Handle other errors here
                        return Promise.reject(error);
                    }
                );

                return instance
            } ()
        },
        reducers: {
            addToken: (state, params) => {
                // state.axios.defaults.headers.common['Authorization'] = `Bearer ${params.payload.token}`;
            },
        }
    }
);

// this is for dispatch
export const {addToken} = applicationSlice.actions;

// this is for configureStore
export default applicationSlice.reducer;