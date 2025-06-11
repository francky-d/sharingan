import { Application } from "../entities/application";

export default interface ApplicationRespositoryInterface {
    getApplications(): Promise<Application[]>
    createApplication(app: Application): Promise<Application>
    updateApplication(app: Application): Promise<Application>
    deleteApplication(ID: Number): Promise<void>
}