namespace go favoritepart

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}


struct FavoriteActionRequest{
    1: string token
    2: i64 video_id
    3: i64 action_type  //1-点赞，2-取消点赞
}
struct FavoriteActionResponse{
    1: i64 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg 
}



service FavoriteService {
    FavoriteActionResponse Favorite(1: FavoriteActionRequest req)
}
