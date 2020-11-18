健身小程序
## 用户需求
由于健身是个长期的过程，我希望有一款软件能帮助我记录在训练中，所做的动作，组数，次数，以及我维度的变化，这样在日积月累过程中，我能用报表或者其他图表类型的数据来展现我的训练成果，或者能帮助我或其他训练者找出更适合自己的健身动作。也能让其他用户分享自己的训练计划。

## 版本计划
1. 所有后端接口
2. 出个h5版本，后期等接口全部完成，再移植到小程序

## 用户场景
1. 在健身过程中，可以记录健身中所做的动作，组数，重量，间歇
2. 以周，月，年形成报表
3. 可以记录自己维度的数据
4. 分享自己的健身计划以及维度
5. 社交属性

## 数据模型

### 用户 （User）
- id
- 用户名  user_name
- 手机号  mobile
- 头像    avatar
- 简介    description
- 微信id  wx_id
- 地区    area
- 年龄    age
- 状态    status
- 创建时间  created_at
- 修改时间  updated_at

### 训练类别(Category)
- id
- 父类id  parent_id
- 类别名称 name
- 类别描述 description
- 类别状态 status
- 创建人   user_id
- 创建时间 created_at
- 修改时间 updated_at


### 训练动作 (Movement)
- id
- 类别 eg：胸,肩,背,腿,手臂,核心，有氧等  cat_id 
- 动作名称 eg: 上斜哑铃卧推    name
- 动作简介                   description
- 状态                      status
- 创建人                     user_id
- 创建时间                   created_at
- 修改时间                   updated_at

### 计划 (Plan)
- id
- 计划名称 plan_name  (eg: 胸部强化训练)
- 状态     status
- 用户id   user_id
- 创建时间  created_at
- 修改时间  updated_at

### 计划详情 (Plan Details)
- id
- plan_id   (属于哪个计划)
- action_id (使用哪个动作)
- 重量        weight
- 次数        count
- 间歇        break
- 状态        status
- 用户        user_id
- 创建时间     created_at
- 修改时间     updated_at

### 训练计划 (Training)

- id    
- 训练日期  training_date
- 使用计划  plan_id
- 开始时间  start_time
- 结束时间  end_time
- 训练小记  description
- 状态      status
- 用户id   user_id  
- 创建时间  created_at
- 修改时间  updated_at

### 训练内容 (Training Logs)

- id
- 计划详情id  plan_detail_id
- 是否完成    done (0 为完成，2, 部分完成  3, 完成 4,失败)
- 状态        status
- 用户        user_id
- 创建时间     created_at
- 修改时间     updated_at

## 接口

### 用户模块
- /api/register [post] 注册用户
- /api/login [post] 登录用户
- /api/users/:id [put] 修改用户信息
- /api/users/forget [post] 忘记密码
- /api/users/reset [post] 重置密码

### 动作类别模块
- /api/categories [get] 获取动作类别
- /api/categories [post] 创建动作类别
- /api/categories/:id [put] 修改单个动作类被
- /api/categories/:id [get] 获取单个动作类别
- /api/categories/id [delete] 删除单个动作类别

### 动作模块
- /api/movements [get]  获取训练动作列表接口
- /api/movements [post] 创建训练动作接口
- /api/movements/:id [put]  修改训练动作接口
- /api/movements/:id [get]  获取单个训练动作接口
- /api/movements/:id [delete]  删除单个训练动作接口

### 计划模块

- /api/plans [post] 创建计划接口
- /api/plans/:id  [get] 获取单个计划接口
- /api/plans/:id [put] 修改计划接口
- /api/plans/:id [delete] 删除计划 

- /api/trainings [post]  新建训练
- /api/trainings/:id/logs [put]  新增一组或多组训练动作到训练计划中
- /api/training/:id/logs [delete] 从训练计划中，删除一组训练动作
- /api/trainings [get] 获取一个人的所有训练计划  未完成
- /api/trainings/:id [put] 修改训练计划  未完成
- /api/trainings/:id [delete] 删除训练计划  未完成

## 页面
### 首页
### 个人中心
制定今日训练计划
查看以往训练计划

