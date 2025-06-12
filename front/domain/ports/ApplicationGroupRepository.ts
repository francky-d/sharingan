import { ApplicationGroup } from "../entities/application-group";

export default interface ApplicationGroupRepositoryInterface {
    getApplicationGroups(): Promise<ApplicationGroup[]>
    getApplicationGroupByID(ID: Number): Promise<ApplicationGroup>
    createApplicationGroup(appGroup: ApplicationGroup): Promise<ApplicationGroup>
    updateApplicationGroup(appGroup: ApplicationGroup): Promise<ApplicationGroup>
    deleteApplicationGroup(ID: Number): Promise<void>
}