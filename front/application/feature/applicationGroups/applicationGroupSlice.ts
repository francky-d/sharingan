import {ApplicationGroup} from "../../../domain/entities/application-group";
import {createSlice} from "@reduxjs/toolkit";
import {fetchAllApplicationsGroupsThunk} from "./applicationGroupThunk";


interface ApplicationGroupState {
    applicationGroups: ApplicationGroup[];
    loading: boolean;
    error: string | null;
}

const initialState: ApplicationGroupState = {
    applicationGroups : [],
    loading: false,
    error: null,
}

const applicationGroupSlice = createSlice({
    name: "applicationGroup",
    initialState,
    reducers: {},
    extraReducers: (builder) =>{
        builder
            .addCase(fetchAllApplicationsGroupsThunk.pending, (state, action) =>{ state.loading = true})
    }
})