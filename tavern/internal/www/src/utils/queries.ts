import { gql } from "@apollo/client";

export const GET_QUEST_QUERY = gql`
    query GetQuests($where: QuestWhereInput) {
        quests(where: $where){
            id
            name
            tasks{
                id
                lastModifiedAt
                output
                execStartedAt
                execFinishedAt
                createdAt
                beacon {
                    id
                    name
                    tags{
                        name
                        kind
                        id   
                    }
                }
            }
            tome{
                id
                name
                paramDefs
            }
        }
    }
`;