/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
    "github.com/golang/glog"
    "github.com/nebulaim/telegramd/proto/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/nebulaim/telegramd/baselib/grpc_util"
    "github.com/nebulaim/telegramd/baselib/logger"
)

// messages.searchStickerSets#c2b7d08b flags:# exclude_featured:flags.0?true q:string hash:int = messages.FoundStickerSets;
func (s *MessagesServiceImpl) MessagesSearchStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("MessagesSearchStickerSets - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl MessagesSearchStickerSets logic

    return nil, fmt.Errorf("Not impl MessagesSearchStickerSets")
}
