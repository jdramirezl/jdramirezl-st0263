

type fileServiceServer struct {
}

func (s *fileServiceServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
        
}

func (s *fileServiceServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
}

func (s *fileServiceServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
}

func (s *fileServiceServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
}
