import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Application } from "../../../domain/entities/application";

import { RootState } from "../../store";
import {fetchApplications} from "./applicationThunk";

interface applicationState {
    applications: Application[];
    loading: boolean;
    error: string | null;
}

const initialState: applicationState = {
    applications: [],
    loading: false,
    error: null,
}

const applicationSlice = createSlice({
    name: "application",
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            //fetch applications
            .addCase(fetchApplications.pending, (state, action) => {
                state.loading = true;
            })
            .addCase(fetchApplications.fulfilled, (state, action) => {
                state.loading=false
                state.applications= action.payload
            })
            .addCase(fetchApplications.rejected, (state, action) => {
               state.loading=false
               state.error = action.error.message
            })
    }
})

export const applicationReducer =  applicationSlice.reducer;
export const {} = applicationSlice.actions;