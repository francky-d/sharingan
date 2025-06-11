import { Application } from "../entities/application";
import ApplicationRepository from "../ports/ApplicationRepository";

export default class ApplicationService {
    constructor(private repository: ApplicationRepository) { }

    async getAllApplications(): Promise<Application[]> {
        return await this.repository.getApplications()
    }

    async createApplication(application: Application) {
        return await this.repository.createApplication(application)
    }

    async updateApplication(application: Application) {
        return await this.repository.updateApplication(application)
    }

    async delelteApplication(ID: Number) {
        return await this.repository.deleteApplication(ID)
    }
}
