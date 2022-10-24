package interactor

// func buildTestModuleTest() *TestUsecase {
// 	var repo repository.DataService
// 	testusecase := NewTestUsecase(repo)
// 	return testusecase
// }

// func TestQueryInfor(t *testing.T) {
// 	type args struct {
// 		ctx    context.Context
// 		testId int
// 		userId int
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    usecase_dto.Test
// 		wantErr error
// 	}{
// 		{
// 			name: "QueryInforSuccess",
// 			args: args{
// 				ctx:    context.Background(),
// 				testId: 1,
// 				userId: 1,
// 			},
// 			want:    usecase_dto.Test{},
// 			wantErr: nil,
// 		},

// 		{
// 			name: "QueryInforError",
// 			args: args{
// 				ctx:    context.Background(),
// 				testId: 1,
// 				userId: 1,
// 			},
// 			want:    usecase_dto.Test{},
// 			wantErr: errors.New("query infor error"),
// 		},
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockTestUseCase(mockCtrl)
// 	testusecase := buildTestModuleTest()

// 	for _, tt := range tests {
// 		if tt.name == "QueryInforSuccess" {
// 			t.Run(tt.name, func(t *testing.T) {
// 				mockRepo.EXPECT().QueryTestInfo(tt.args.ctx, tt.args.testId, tt.args.userId).Return(tt.want, tt.wantErr)
// 				got, err := testusecase.QueryTestInfo(tt.args.ctx, tt.args.testId, tt.args.userId)
// 				if err != nil {
// 					t.Errorf("TestUsecase.QueryInfor() error = %v", err)
// 				}
// 				if !reflect.DeepEqual(got, tt.want) {
// 					t.Errorf("TestUsecase.QueryInfor() = %v, want %v", got, tt.want)
// 				}
// 			})
// 		}
// 		if tt.name == "QueryInforError" {
// 			t.Run(tt.name, func(t *testing.T) {
// 				mockRepo.EXPECT().QueryTestInfo(tt.args.ctx, tt.args.testId, tt.args.userId).Return(tt.want, tt.wantErr)
// 				got, err := testusecase.QueryTestInfo(tt.args.ctx, tt.args.testId, tt.args.userId)
// 				if err == nil {
// 					t.Errorf("TestUsecase.QueryInfor() error = %v", err)
// 				}
// 				if !reflect.DeepEqual(got, tt.want) {
// 					t.Errorf("TestUsecase.QueryInfor() = %v, want %v", got, tt.want)
// 				}
// 			})
// 		}
// 	}
// }

// func TestQuerySkillTest(t *testing.T) {
// 	type args struct {
// 		ctx    context.Context
// 		testId int
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    usecase_dto.Test
// 		wantErr error
// 	}{
// 		{
// 			name: "QuerySkillTest",
// 			args: args{
// 				ctx:    context.Background(),
// 				testId: 1,
// 			},
// 			want:    usecase_dto.Test{},
// 			wantErr: nil,
// 		},
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockTestUseCase(mockCtrl)
// 	testusecase := buildTestModuleTest()

// 	for _, tt := range tests {
// 		if tt.name == "QuerySkillTest" {
// 			t.Run(tt.name, func(t *testing.T) {
// 				mockRepo.EXPECT().QuerySkillTest(tt.args.ctx, tt.args.testId).Return(tt.want, tt.wantErr)
// 				got, err := testusecase.QuerySkillTest(tt.args.ctx, tt.args.testId)
// 				if err != nil {
// 					t.Errorf("TestUsecase.QuerySkillTest() error = %v", err)
// 				}
// 				if !reflect.DeepEqual(got, tt.want) {
// 					t.Errorf("TestUsecase.QuerySkillTest() = %v, want %v", got, tt.want)
// 				}
// 			})
// 		}
// 	}
// }

// func TestSubmitTest(t *testing.T) {
// 	type args struct {
// 		ctx        context.Context
// 		data       usecase_dto.SubmitData
// 		userId     int
// 		entityCode int
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    usecase_dto.Test
// 		wantErr error
// 	}{
// 		{
// 			name: "SubmitTest",
// 			args: args{
// 				ctx:    context.Background(),
// 				data:   usecase_dto.SubmitData{},
// 				userId: 1,
// 			},
// 			want:    usecase_dto.Test{},
// 			wantErr: nil,
// 		},
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockTestUseCase(mockCtrl)
// 	testusecase := buildTestModuleTest()

// 	for _, tt := range tests {
// 		if tt.name == "SubmitTest" {
// 			t.Run(tt.name, func(t *testing.T) {
// 				mockRepo.EXPECT().SubmitTest(tt.args.ctx, tt.args.data, tt.args.userId, tt.args.entityCode).Return(tt.want, tt.wantErr)
// 				got, err := testusecase.SubmitTest(tt.args.ctx, tt.args.data, tt.args.userId, tt.args.entityCode)
// 				if err != nil {
// 					t.Errorf("TestUsecase.SubmitTest() error = %v", err)
// 				}
// 				if !reflect.DeepEqual(got, tt.want) {
// 					t.Errorf("TestUsecase.SubmitTest() = %v, want %v", got, tt.want)
// 				}
// 			})
// 		}
// 	}
// }

// func TestQueryAllTest(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockTestUseCase(mockCtrl)
// 	testusecase := buildTestModuleTest()

// 	t.Run("Query All Test", func(t *testing.T) {
// 		mockRepo.EXPECT().QueryAllTest(context.Background()).Return([]usecase_dto.Test{}, nil)
// 		got, err := testusecase.QueryAllTest(context.Background())
// 		if err != nil {
// 			t.Errorf("TestUsecase.QueryAllTest() error = %v", err)
// 		}
// 		if !reflect.DeepEqual(got, []usecase_dto.Test{}) {
// 			t.Errorf("TestUsecase.QueryAllTest() = %v, want %v", got, []usecase_dto.Test{})
// 		}
// 	})
// }

// func QueryTestAnswer(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockTestUseCase(mockCtrl)
// 	testusecase := buildTestModuleTest()

// 	t.Run("Query Test Answer", func(t *testing.T) {
// 		mockRepo.EXPECT().QueryTestAnswer(context.Background(), 1).Return([]usecase_dto.SubmitData{}, nil)
// 		got, err := testusecase.QueryTestAnswer(context.Background(), 1)
// 		if err != nil {
// 			t.Errorf("TestUsecase.QueryTestAnswer() error = %v", err)
// 		}
// 		if !reflect.DeepEqual(got, []usecase_dto.SubmitData{}) {
// 			t.Errorf("TestUsecase.QueryTestAnswer() = %v, want %v", got, []usecase_dto.SubmitData{})
// 		}
// 	})
// }
