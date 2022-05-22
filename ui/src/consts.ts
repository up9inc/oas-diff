export enum ChangeTypeEnum {
    Created = "create",
    Updated = "update",
    Deleted = "delete"
}

export const createClass = ChangeTypeEnum.Created
export const updateClass = ChangeTypeEnum.Updated
export const deleteClass = ChangeTypeEnum.Deleted
export const infoClass = "info"
