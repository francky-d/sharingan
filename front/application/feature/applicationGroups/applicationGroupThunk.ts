import {createAsyncThunk} from "@reduxjs/toolkit";
import ApplicationGroupService from "../../../domain/services/applicationGroupService";
import ApplicationApi from "../../../infrastructure/api/applicationApi";
import ApplicationGroupApi from "../../../infrastructure/api/ApplicationGroupApi";

const applicationGrpRepository =  new ApplicationGroupApi()
const applicationGrpService = new ApplicationGroupService(applicationGrpRepository)

const  fetchAllApplicationsGroupsThunk =  createAsyncThunk(
    "applicationGroups/fetchAllApplicationsGroups",
         (payload, thunkAPI) => {
            return applicationGrpService.getApplicationGroups();
        }
    )

export {fetchAllApplicationsGroupsThunk}