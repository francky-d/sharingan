import ApplicationGroupRepositoryInterface from "../ports/ApplicationGroupRepository";
import {ApplicationGroup} from "../entities/application-group";


export default class ApplicationGroupService {
    constructor(private readonly repository: ApplicationGroupRepositoryInterface){}

    getApplicationGroups(): Promise<ApplicationGroup[]>{
        return this.repository.getApplicationGroups();
    }

    getApplicationGroupsByID(ID: number): Promise<ApplicationGroup>{
       return this.repository.getApplicationGroupByID(ID)
    }
}