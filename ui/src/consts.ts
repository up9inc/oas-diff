export enum ChangeTypeEnum {
    Created = "create",
    Updated = "update",
    Deleted = "delete"
}

export const createClass = ChangeTypeEnum.Created
export const updateClass = ChangeTypeEnum.Updated
export const deleteClass = ChangeTypeEnum.Deleted
export const infoClass = "info"

type EnumDictionary<T extends string | symbol | number, U> = {
    [K in T]: U;
};
export const TypeCaptionDictionary: EnumDictionary<string, string> = {
    [ChangeTypeEnum.Created]: "Created",
    [ChangeTypeEnum.Updated]: "Updated",
    [ChangeTypeEnum.Deleted]: "Deleted"
}
