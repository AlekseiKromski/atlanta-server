import {createSlice} from '@reduxjs/toolkit';
import axios from "axios"

const applicationSlice = createSlice({
        name: 'axios',
        initialState: {
            axios: axios.create({
                baseURL: process.env.REACT_APP_AXIOS_BASE_URL,
                timeout: 1000,
                headers: {'X-Custom-Header': 'foobar'}
            })
        },
        reducers: {}
    }
);

// this is for dispatch
// export const {addTodo} = applicationSlice.actions;

// this is for configureStore
export default applicationSlice.reducer;