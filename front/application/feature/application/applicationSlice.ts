import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Application } from "../../../domain/entities/application";
import { fetchApplications } from "./applicationsThunks";
import { RootState } from "../../store";

interface applicationState {
    applications: []
}

const initialState: applicationState = {
    applications: [],
}

const applicationSlice = createSlice({
    name: "application",
    initialState,
    reducers: {
        addApplication(state: RootState, action: PayloadAction<Application>) {

        }
    }
})