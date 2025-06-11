import { createAsyncThunk } from "@reduxjs/toolkit";
import { Application } from "../../../domain/entities/application";
import ApplicationService from "../../../domain/services/applicationService";
import ApplicationApi from "../../../infrastructure/api/applicationApi";

const applicationRepository = new ApplicationApi()
const applicationService = new ApplicationService(applicationRepository)

export const fetchApplications = createAsyncThunk("applications/fetchApplications",
    async () => { return applicationService.getAllApplications() }
)

export const createApplication = createAsyncThunk("applications/createApplicaton",
    async (applicaton: Application) => {
        return await applicationService.createApplication(applicaton)
    }
)