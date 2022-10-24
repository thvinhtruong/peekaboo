package interactor

// func buildClassTest() *ClassUsecase {
// 	var repo repository.DataService
// 	classusecase := NewClassUsecase(repo)
// 	return classusecase

// }

// func TesrCreateClass(t *testing.T) {
// 	type args struct {
// 		ctx       context.Context
// 		class_dto usecase_dto.Class
// 	}

// 	type test struct {
// 		name string
// 		args args
// 		want error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tests := []test{
// 		{
// 			name: "TestCreateClassSuccess",
// 			args: args{
// 				ctx:       context.Background(),
// 				class_dto: usecase_dto.Class{},
// 			},
// 			want: nil,
// 		},

// 		{
// 			name: "TestCreateClassError",
// 			args: args{
// 				ctx:       context.Background(),
// 				class_dto: usecase_dto.Class{},
// 			},
// 			want: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.name == "TestCreateClassError" {
// 				mockRepo.EXPECT().CreateClass(tt.args.ctx, tt.args.class_dto).Return(errors.New("error when creating class"))
// 				classusecase := buildClassTest()
// 				if got := classusecase.CreateClass(tt.args.ctx, tt.args.class_dto); got != tt.want {
// 					t.Errorf("ClassUsecase.CreateClass() = %v, want %v", got, tt.want)
// 				}
// 			}
// 			if tt.name == "TestCreateClassSuccess" {
// 				mockRepo.EXPECT().CreateClass(tt.args.ctx, tt.args.class_dto).Return(nil)
// 				classusecase := buildClassTest()
// 				if got := classusecase.CreateClass(tt.args.ctx, tt.args.class_dto); got != tt.want {
// 					t.Errorf("ClassUsecase.CreateClass() = %v, want %v", got, tt.want)
// 				}
// 			}
// 		})
// 	}

// }

// func TestGetClasses(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}

// 	type test struct {
// 		name    string
// 		args    args
// 		want    []usecase_dto.Class
// 		wantErr error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tt := test{
// 		name: "TestGetClasses",
// 		args: args{
// 			ctx: context.Background(),
// 		},
// 		want:    []usecase_dto.Class{},
// 		wantErr: nil,
// 	}

// 	t.Run(tt.name, func(t *testing.T) {
// 		mockRepo.EXPECT().GetClasses(tt.args.ctx).Return(tt.want, tt.wantErr)
// 		classusecase := buildClassTest()
// 		got, err := classusecase.GetClasses(tt.args.ctx)
// 		if err != nil {
// 			t.Errorf("ClassUsecase.GetClasses() = %v, want %v", err, tt.wantErr)
// 		}
// 		if !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("ClassUsecase.GetClasses() = %v, want %v", got, tt.want)
// 		}

// 	})
// }

// func TestAddMember2Class(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		classId int
// 		userId  int
// 	}

// 	type test struct {
// 		name string
// 		args args
// 		want error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tt := test{
// 		name: "TestAddMember2Class",
// 		args: args{
// 			ctx:     context.Background(),
// 			classId: 1,
// 			userId:  1,
// 		},
// 		want: nil,
// 	}

// 	t.Run(tt.name, func(t *testing.T) {
// 		mockRepo.EXPECT().AddMember2Class(tt.args.ctx, tt.args.classId, tt.args.userId).Return(tt.want)
// 		classusecase := buildClassTest()
// 		if got := classusecase.AddMember2Class(tt.args.ctx, tt.args.classId, tt.args.userId); got != tt.want {
// 			t.Errorf("ClassUsecase.AddMember2Class() = %v, want %v", got, tt.want)
// 		}
// 	})
// }

// func TestRemoveMemberFromClass(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		classId int
// 		userId  int
// 	}

// 	type test struct {
// 		name string
// 		args args
// 		want error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tt := test{
// 		name: "TestRemoveMemberFromClass",
// 		args: args{
// 			ctx:     context.Background(),
// 			classId: 1,
// 			userId:  1,
// 		},
// 		want: nil,
// 	}

// 	t.Run(tt.name, func(t *testing.T) {
// 		mockRepo.EXPECT().RemoveMemberFromClass(tt.args.ctx, tt.args.classId, tt.args.userId).Return(tt.want)
// 		classusecase := buildClassTest()
// 		if got := classusecase.RemoveMemberFromClass(tt.args.ctx, tt.args.classId, tt.args.userId); got != tt.want {
// 			t.Errorf("ClassUsecase.RemoveMemberFromClass() = %v, want %v", got, tt.want)
// 		}
// 	})
// }

// func TestQueryClassTestResult(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		classId int
// 		testId  int
// 	}

// 	type test struct {
// 		name    string
// 		args    args
// 		want    usecase_dto.TestResult
// 		wantErr error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tt := test{
// 		name: "TestQueryClassTestResult",
// 		args: args{
// 			ctx:     context.Background(),
// 			classId: 1,
// 			testId:  1,
// 		},
// 		want:    usecase_dto.TestResult{},
// 		wantErr: nil,
// 	}

// 	t.Run(tt.name, func(t *testing.T) {
// 		mockRepo.EXPECT().QueryClassTestResult(tt.args.ctx, tt.args.classId, tt.args.testId).Return(tt.want, tt.wantErr)
// 		classusecase := buildClassTest()
// 		got, err := classusecase.QueryClassTestResult(tt.args.ctx, tt.args.classId, tt.args.testId)
// 		if err != nil {
// 			t.Errorf("ClassUsecase.QueryClassTestResult() = %v, want %v", err, tt.wantErr)
// 		}
// 		if !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("ClassUsecase.QueryClassTestResult() = %v, want %v", got, tt.want)
// 		}
// 	})
// }

// func TestGetClassTest(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		classId int
// 	}

// 	type test struct {
// 		name    string
// 		args    args
// 		want    usecase_dto.Test
// 		wantErr error
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRepo := mockusecase.NewMockClassUseCase(mockCtrl)

// 	tt := test{
// 		name: "TestGetClassTest",
// 		args: args{
// 			ctx:     context.Background(),
// 			classId: 1,
// 		},
// 	}

// 	t.Run(tt.name, func(t *testing.T) {
// 		mockRepo.EXPECT().GetClassTest(tt.args.ctx, tt.args.classId).Return(tt.want, tt.wantErr)
// 		classusecase := buildClassTest()
// 		got, err := classusecase.GetClassTest(tt.args.ctx, tt.args.classId)
// 		if err != nil {
// 			t.Errorf("ClassUsecase.GetClassTest() = %v, want %v", err, tt.wantErr)
// 		}
// 		if !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("ClassUsecase.GetClassTest() = %v, want %v", got, tt.want)
// 		}
// 	})
// }
