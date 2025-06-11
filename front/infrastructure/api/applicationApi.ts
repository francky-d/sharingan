import { Application } from "../../domain/entities/application";
import ApplicationRespositoryInterface from "../../domain/ports/ApplicationRepository";
import axios from "axios"

export default class ApplicationApi implements ApplicationRespositoryInterface {
    private apiUrl = "/api/v1/applications"


    async getApplications(): Promise<Application[]> {
        const response = await axios.get<Application[]>(this.apiUrl)
        return response.data
    }

    async createApplication(application: Application): Promise<Application> {
        const response = await axios.post<Application>(this.apiUrl, application)
        return response.data
    }

    async updateApplication(application: Application): Promise<Application> {
        const response = await axios.put<Application>(this.apiUrl, application)
        return response.data

    }

    async deleteApplication(ID: Number): Promise<void> {
        await axios.delete(`${this.apiUrl}/${ID}`)
    }



}